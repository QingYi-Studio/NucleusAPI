# NucleusAPI

NucleusAPI is a comprehensive database and API for computer hardware components, providing detailed specifications for CPUs, APUs, GPUs, NPUs, and SoCs from various manufacturers.

## Features

- **Extensive Database**: Contains detailed JSON files for a wide range of hardware components.
- **API Endpoints**: Offers API routes to access the data programmatically.
- **Manufacturer Support**: Includes data from AMD, Apple, Intel, and other manufacturers.
- **Hardware Types**: Covers CPUs, APUs, GPUs, NPUs, and SoCs.

## Directory Structure

The project is organized into several directories:

- `data`: Contains JSON files with hardware specifications.
  - `apu`: Data for Accelerated Processing Units (APUs).
  - `cpu`: Data for Central Processing Units (CPUs).
  - `gpu`: Data for Graphics Processing Units (GPUs).
  - `npu`: Data for Neural Processing Units (NPUs).
  - `soc`: Data for System on Chips (SoCs).
- `handlers`: Go files containing handlers for API requests.
- `pages`: (Directory purpose not specified in the given structure).
- `routes`: Go files defining the API routes.

## Example JSON Schema

Each JSON file follows a specific schema. Here is an example schema for an APU:

```json
{
  "name": "AMD A12-9800",
  "codename": "...",
  "architecture": "...",
  "process_node": {
    "cpu_cores": "...",
    "io_die": "..."
  },
  "transistors": ...,
  "die_size_mm2": 250,
  "package": "...",
  "launch_date": "...",
  "part_number": "...",
  "launch_price_usd": ...,
  "cores": {
    "total": ...,
    "performance_cores": ...,
    "efficient_cores": ...,
    "threads": ...
  },
  "clock": {
    "base_ghz": ...,
    "max_boost_ghz": ...,
    "base_clock_mhz": ...,
    "multiplier": ...,
    "unlocked_multiplier": ...
  },
  "cache": {
    "l1_instruction_kb_per_module": ...,
    "l1_data_kb_per_core": ...,
    "l1_total_kb_per_module": ...,
    "l2_mb_total": ...,
    "l2_mb_per_module": ...,
    "l3_mb_total": ...
  },
  "tdp": {
    "tdp_w": ...,
    "ppt_w": ...
  },
  "temperature": {
    "tjunction_max_c": ...,
    "tcase_max_c": ...
  },
  "memory": {
    "type": ...,
    "native_speed_mt_s": ...,
    "max_overclock_mt_s": ...,
    "channels": ...,
    "max_capacity_gb": ...,
    "ecc": ...
  },
  "pcie": {
    "version": ...,
    "lanes": ...
  },
  "integrated_graphics": {...},
  "features": [...],
  "chipset_compatibility": [...],
  "cooler_included": ...,
  "recommended_cooler": ...
}
```

## Getting Started

To use NucleusAPI, you can either access the data directly from the `data` directory or use the provided API endpoints.

### Accessing Data

Navigate to the `data` directory and choose the relevant subdirectory (e.g., `cpu`, `apu`) to find the JSON files with hardware specifications.

### Using the API

The API is defined in the `routes` directory. To start the server and access the API, run the main Go application.

## Contribution

Contributions to NucleusAPI are welcome. Please ensure that any new data files adhere to the established JSON schema.

## License

This project is open-source and available under the Apache License 2.0. See the LICENSE file for more details.

---

Please note that this README is a template and may require adjustments to accurately reflect the actual structure and functionality of your project.
