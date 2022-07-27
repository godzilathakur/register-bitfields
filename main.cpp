#include <iostream>
#include "register_api.h"

using namespace std;

int main() {
    register_example_api_t reg{};
    auto calibrate = reg.read_calibrate();
    cout << static_cast<unsigned int>(calibrate) << endl;
    reg.write_calibrate(register_example_api_t::calibrate_t::CALIBRATE);
    cout << static_cast<unsigned int>(reg.read_calibrate()) << endl;

    return 0;
}
