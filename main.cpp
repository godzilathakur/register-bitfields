#include <iostream>
#include "fancy_register_defs.h"

using namespace std;

int main(void) {
    init_block_t init_block = {0};

    init_block.status = init_block_t::INITIALIZED;
    cout<< "status " << init_block.status << " reset " << init_block.reset << endl;
    init_block.reset = 1;
    init_block.status = init_block_t::RESET;
    cout<< "status " << init_block.status << " reset " << init_block.reset << endl;
}