#pragma once
// auto-generated file

struct init_block_t {
  unsigned int status : 3;
  enum status_t {
    WATCHDOG = 5,
    INITIALIZED = 6,
    READY = 7,
    UNINITIALIZED = 0,
    RESET = 1,
    CALIBRATION = 2,
    RESERVED = 3,
    OVERCURRENT = 4,
  };
  unsigned int mode : 2;
  unsigned int reset : 1;
};

struct command_t {
  unsigned int module : 3;
  unsigned int assert : 1;
};
