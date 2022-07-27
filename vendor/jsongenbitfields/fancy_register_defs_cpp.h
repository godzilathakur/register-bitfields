#pragma once
// auto-generated file using bitfield-gen

struct init_block_t {
  unsigned int status : 3;
  enum class status_t {
    INITIALIZED = 6,
    READY = 7,
    UNINITIALIZED = 0,
    RESET = 1,
    CALIBRATION = 2,
    RESERVED = 3,
    OVERCURRENT = 4,
    WATCHDOG = 5,
  };
  unsigned int mode : 2;
  enum class mode_t {
    FREERUNNING = 0,
    STEPPED = 1,
    RESERVED = 2,
    DISABLED = 3,
  };
  unsigned int reset : 1;
};

struct command_t {
  unsigned int module : 3;
  enum class module_t {
    ACTUATOR = 2,
    GPIO = 4,
  };
  unsigned int assert : 1;
};

