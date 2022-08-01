#include <iostream>
#include "register_api.h"
#include "multiplexer_api.h"

using namespace std;

int main() {
    register_example_api_t reg{};
    auto calibrate = reg.read_calibrate();
    cout << static_cast<unsigned int>(calibrate) << endl;
    reg.write_calibrate(register_example_api_t::calibrate_t::CALIBRATE);
    cout << static_cast<unsigned int>(reg.read_calibrate()) << endl;

    register_control_reg_t contol_reg{};
    auto selection = contol_reg.read_channel_selection();
    cout << "selection before: " << static_cast<unsigned int>(selection) << endl;

    contol_reg.write_channel_selection(register_control_reg_t::channel_selection_t::CHANNEL_1);
    selection = contol_reg.read_channel_selection();
    cout << "selection after: " << static_cast<unsigned int>(selection) << endl;

    auto enable = contol_reg.read_enable();
    cout << "enable before: " << static_cast<unsigned int>(enable) << endl;

    contol_reg.write_enable(1);
    enable = contol_reg.read_enable();
    cout << "enable after: " << static_cast<unsigned int>(enable) << endl;

    return 0;
}
