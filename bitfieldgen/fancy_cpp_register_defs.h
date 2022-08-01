#pragma once
/* auto-generated file using bitfieldgen

 Peripheral Name Fancy
 Description 
 Specifications 
*/

constexpr static unsigned int c_peripheral_addr = 0;
void io_write(const unsigned int reg_addr, const unsigned int data) {

	};

void io_read(const unsigned int reg_addr, unsigned int& data) {

	};

class register_init_block_t {
private:
  constexpr static unsigned int c_register_init_block_addr = 0;
  union register_defs_t {
    struct fields_t {
      unsigned int status : 3; // READ_ONLY
      unsigned int reserved_0 : 2; // UNSUPPORTED
      unsigned int mode : 2; // READ_WRITE
      unsigned int reset : 1; // WRITE_ONLY
    } m_fields;
    unsigned int m_data;

    register_defs_t() {
		io_read(c_register_init_block_addr, m_data);
	};
    register_defs_t& operator=(const register_defs_t& other) {
		m_data = other.m_data;
		return *this;
	};
    void operator=(const unsigned int val) {
		m_data = val;
		io_write(c_register_init_block_addr, m_data);
	};
  };

public:
  enum class status_t {
    RESERVED = 3,
    OVERCURRENT = 4,
    WATCHDOG = 5,
    INITIALIZED = 6,
    READY = 7,
    UNINITIALIZED = 0,
    RESET = 1,
    CALIBRATION = 2,
  };

  enum class mode_t {
    FREERUNNING = 0,
    STEPPED = 1,
    RESERVED = 2,
    DISABLED = 3,
  };

  register_init_block_t() = default;
	~register_init_block_t() = default;
  status_t read_status() const {
    auto defs = register_defs_t{};
    return static_cast<status_t>(defs.m_fields.status);
  };

  mode_t read_mode() const {
    auto defs = register_defs_t{};
    return static_cast<mode_t>(defs.m_fields.mode);
  };

  void write_mode(mode_t val) {
    auto defs = register_defs_t{};
    defs.m_fields.mode = static_cast<unsigned int>(val);
    defs = defs.m_data;
  };

  void write_reset(unsigned int val) {
    auto defs = register_defs_t{};
    defs.m_fields.reset = static_cast<unsigned int>(val);
    defs = defs.m_data;
  };

};


class register_command_t {
private:
  constexpr static unsigned int c_register_command_addr = 0;
  union register_defs_t {
    struct fields_t {
      unsigned int module : 3; // READ_WRITE
      unsigned int assert : 1; // WRITE_ONLY
    } m_fields;
    unsigned int m_data;

    register_defs_t() {
		io_read(c_register_command_addr, m_data);
	};
    register_defs_t& operator=(const register_defs_t& other) {
		m_data = other.m_data;
		return *this;
	};
    void operator=(const unsigned int val) {
		m_data = val;
		io_write(c_register_command_addr, m_data);
	};
  };

public:
  enum class module_t {
    GPIO = 4,
    ACTUATOR = 2,
  };

  register_command_t() = default;
	~register_command_t() = default;
  module_t read_module() const {
    auto defs = register_defs_t{};
    return static_cast<module_t>(defs.m_fields.module);
  };

  void write_module(module_t val) {
    auto defs = register_defs_t{};
    defs.m_fields.module = static_cast<unsigned int>(val);
    defs = defs.m_data;
  };

  void write_assert(unsigned int val) {
    auto defs = register_defs_t{};
    defs.m_fields.assert = static_cast<unsigned int>(val);
    defs = defs.m_data;
  };

};


