#include <iostream>
using namespace std;

void io_write(const unsigned int data) {

};
void io_read(unsigned int& data) {

};

class register_example_api_t {
private:
    union register_defs_t {
        struct fields_t {
            unsigned int status : 3; // read only
            unsigned int reserved : 3; // no read or write
            unsigned int calibrate : 1; // read-write
            unsigned int reset : 1; // write only
        } m_fields;
        unsigned int m_data;

        register_defs_t() {
            io_read(m_data);
        }

        register_defs_t& operator=(const register_defs_t& other) {
            m_data = other.m_data;
            return *this;
        }

        void operator=(const unsigned int val) {
            m_data = val;
            io_write(m_data);
        }
    };

public:
    enum class status_t {
        INIT = 0,
        COMMAND = 1,
        DEBUG = 2,
        RESERVED = 3,
        TEMP = 4,
        OTP = 5
    };
    enum class calibrate_t {
        INACTIVE = 0,
        CALIBRATE = 1
    };
    enum class reset_t {
        INACTIVE = 0,
        RESET = 1
    };

    register_example_api_t() = default;
    ~register_example_api_t() = default;

    status_t read_status() const {
        auto defs = register_defs_t{};
        return static_cast<status_t>(defs.m_fields.status);
    }
    calibrate_t read_calibrate() const {
        auto defs = register_defs_t{};
        return static_cast<calibrate_t>(defs.m_fields.calibrate);
    }
    void write_calibrate(calibrate_t val) {
        auto defs = register_defs_t{};
        defs.m_fields.calibrate = static_cast<unsigned int>(val);
        defs = defs.m_data;
    }
    void write_reset(reset_t val) {
        auto defs = register_defs_t{};
        defs.m_fields.reset = static_cast<unsigned int>(val);
        defs = defs.m_data;
    }
};