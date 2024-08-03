#include <iostream>
#include <omp.h>
#include <cstdlib>
#include <ctime>

using namespace std;

bool isInside(double x, double y);
double calculateDistance(double x, double y);

int main()
{
    int iterations;
    cout << "Please input the number of iterations: ";
    cin >> iterations;

    // Seed the random number generator
    srand(time(0));

    // start time
    double start = omp_get_wtime();

    long double insidee = 0;

#pragma omp parallel reduction(+:insidee)
    {
        unsigned int seed = omp_get_thread_num();
        #pragma omp for
        for (int i = 0; i < iterations; i++)
        {
            double x = (double)rand_r(&seed) / RAND_MAX;
            double y = (double)rand_r(&seed) / RAND_MAX;

            if (isInside(x, y))
            {
                insidee++;
            }
        }
    }

    // end time
    double end = omp_get_wtime();

    long double pi = 4.0 * insidee / iterations;
    cout << "Pi is approximately " << pi << endl;
    cout << "Time: " << end - start << "s" << endl;

    return 0;
}

double calculateDistance(double x, double y)
{
    return x * x + y * y;
}

bool isInside(double x, double y)
{
    return calculateDistance(x, y) <= 1.0;
}
