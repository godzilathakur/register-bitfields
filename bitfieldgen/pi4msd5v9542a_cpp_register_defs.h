#pragma once
/* auto-generated file using bitfieldgen

 Peripheral Name PI4MSD5V9542A
 Description 2 Channel I2C bus Multiplexer
 Specifications https://www.diodes.com/assets/Datasheets/PI4MSD5V9542A.pdf
*/

constexpr static unsigned int c_peripheral_addr = 64;
void io_write(const unsigned int reg_addr, const unsigned int data) {

	};

void io_read(const unsigned int reg_addr, unsigned int& data) {

	};

class register_control_reg_t {
private:
  constexpr static unsigned int c_register_control_reg_addr = 0;
  union register_defs_t {
    struct fields_t {
      unsigned int interrupts : 2; // READ_ONLY
      unsigned int reserved_0 : 1; // UNSUPPORTED
      unsigned int enable : 1; // READ_WRITE
      unsigned int channel_selection : 2; // READ_WRITE
    } m_fields;
    unsigned int m_data;

    register_defs_t() {
		io_read(c_register_control_reg_addr, m_data);
	};
    register_defs_t& operator=(const register_defs_t& other) {
		m_data = other.m_data;
		return *this;
	};
    void operator=(const unsigned int val) {
		m_data = val;
		io_write(c_register_control_reg_addr, m_data);
	};
  };

public:
  enum class interrupts_t {
    INT_0 = 0,
    INT_1 = 1,
  };

  enum class channel_selection_t {
    CHANNEL_0 = 0,
    CHANNEL_1 = 1,
  };

  register_control_reg_t() = default;
	~register_control_reg_t() = default;
  interrupts_t read_interrupts() const {
    auto defs = register_defs_t{};
    return static_cast<interrupts_t>(defs.m_fields.interrupts);
  };

  unsigned int read_enable() const {
    auto defs = register_defs_t{};
    return defs.m_fields.enable;
  };

  void write_enable(unsigned int val) {
    auto defs = register_defs_t{};
    defs.m_fields.enable = static_cast<unsigned int>(val);
    defs = defs.m_data;
  };

  channel_selection_t read_channel_selection() const {
    auto defs = register_defs_t{};
    return static_cast<channel_selection_t>(defs.m_fields.channel_selection);
  };

  void write_channel_selection(channel_selection_t val) {
    auto defs = register_defs_t{};
    defs.m_fields.channel_selection = static_cast<unsigned int>(val);
    defs = defs.m_data;
  };

};


