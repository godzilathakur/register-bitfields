BitField Gen
============

Generate C/C++ Struct BitFields based on register specification in JSON

Struct bitfields C/C++ is a popular way of encapsulating peripheral registers.
They facilitate  efficient memory usage, as well as provide read-modify-write
without need for bitmasks.

This program generates struct bitfields by parsing register definitions.

In order to describe the register definitions, we use a JSON format as follows

```
 {
   "peripheral_name": # string - name of the peripheral
   "peripheral_address": # int - I2C slave address
  "config": {
    "width": # float64 - width of each register in bits, default: 8
  },
  "registers": [
    {
      "name":  # string - name of the register
      "address": # int - I2C register address
      
      "fields": [
        {
          "name": # string - name of the register field

          "attribute": # string - RO read-only, W write-only or RW read-write field. 
          Reserved registers do not have an attribute
          
          // bit field range is defined using msb and lsb as follows:
          
          "msb": # float64 - highest bit needed by field
          
          "lsb": # float64 - lowest bit needed by field. LSB = MSB for a single bit field
          
          "values": {
			/* possible field values that fit the 
			aforementioned bit field range. 
			these will be converted to enum definition */
			
            # name: value,
          }
        }
      ]
    }
  ]
 }

```

Please refer the example json definition files

## Usage 
Define definitions.json file per the register spec you need. 

```shell
cd bitfieldgen
go build .
./bitfieldgen -file <your_filename>.json -gencpp
```
The program generates C++ header file in the format: <your_peripheral_name>.h
