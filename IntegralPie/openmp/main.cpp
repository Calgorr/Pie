#include <iostream>
#include <omp.h>

using namespace std;

int main()
{
    long long int n;
    cout << "Please input the number of iterations: ";
    cin >> n;
    double sum = 0.0;
    double step = 1.0 / n;
    double x;

    double start = omp_get_wtime();

#pragma omp parallel for reduction(+ : sum) private(x)
    for (long long int i = 0; i < n; i++)
    {
        x = (i + 0.5) * step;
        sum += 4.0 / (1.0 + x * x);
    }

    double pi = step * sum;
    double end = omp_get_wtime();

    cout << "Pi is approximately " << pi << endl;
    cout << "Time: " << end - start << "s" << endl;

    return 0;
}