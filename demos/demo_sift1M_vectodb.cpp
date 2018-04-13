#include "vectodb.h"
#include <iostream>
#include <memory>

#include <string.h>
#include <sys/stat.h>
#include <sys/time.h>
#include <sys/types.h>
#include <unistd.h>

#include <cassert>

using namespace std;

/**
 * To run this demo, please download the ANN_SIFT1M dataset from
 *
 *   http://corpus-texmex.irisa.fr/
 *
 * and unzip it to the sudirectory sift1M.
 * 
 * This demo trains an index for the given database.
 **/

/*****************************************************
 * I/O functions for fvecs and ivecs
 *****************************************************/

float*
fvecs_read(const char* fname, size_t* d_out, size_t* n_out)
{
    FILE* f = fopen(fname, "r");
    if (!f) {
        fprintf(stderr, "could not open %s\n", fname);
        perror("");
        abort();
    }
    int d;
    fread(&d, 1, sizeof(int), f);
    assert((d > 0 && d < 1000000) || !"unreasonable dimension");
    fseek(f, 0, SEEK_SET);
    struct stat st;
    fstat(fileno(f), &st);
    size_t sz = st.st_size;
    assert(sz % ((d + 1) * 4) == 0 || !"weird file size");
    size_t n = sz / ((d + 1) * 4);

    *d_out = d;
    *n_out = n;
    float* x = new float[n * (d + 1)];
    size_t nr = fread(x, sizeof(float), n * (d + 1), f);
    assert(nr == n * (d + 1) || !"could not read whole file");

    // shift array to remove row headers
    for (size_t i = 0; i < n; i++)
        memmove(x + i * d, x + 1 + i * (d + 1), d * sizeof(*x));

    fclose(f);
    return x;
}

// not very clean, but works as long as sizeof(int) == sizeof(float)
int* ivecs_read(const char* fname, size_t* d_out, size_t* n_out)
{
    return (int*)fvecs_read(fname, d_out, n_out);
}

double
elapsed()
{
    struct timeval tv;
    gettimeofday(&tv, nullptr);
    return tv.tv_sec + tv.tv_usec * 1e-6;
}

// train phase, input: index_key database train_set, output: index
int main()
{
    double t0 = elapsed();

    printf("[%.3f s] Loading database\n", elapsed() - t0);
    //auto vdb{ std::make_unique<VectoDB>("/tmp", 128, 1) };
    VectoDB vdb("/tmp", 128, 1);
    size_t nb, d;
    float* xb = fvecs_read("sift1M/sift_learn.fvecs", &d, &nb);
    long* xids = new long[nb];
    for (long i = 0; i < (long)nb; i++) {
        xids[i] = i;
    }
    vdb.AddWithIds(nb, xb, xids);
    delete[] xb;
    delete[] xids;

    printf("[%.3f s] Building index\n", elapsed() - t0);
    faiss::Index* index;
    vdb.BuildIndex(index);
    vdb.ActivateIndex(index);

    printf("[%.3f s] Searching index\n", elapsed() - t0);
    size_t nq;
    size_t d2;
    float* xq = fvecs_read("sift1M/sift_query.fvecs", &d2, &nq);
    float* D = new float[nq];
    long* I = new long[nq];
    vdb.Search(nq, xq, D, I);

    printf("[%.3f s] Compute recalls\n", elapsed() - t0);
    delete[] D;
    delete[] I;
}