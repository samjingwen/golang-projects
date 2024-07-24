# CPX CLI

## Overview

CPX CLI is a command-line tool to manage and monitor cloud services deployed on CPX. It provides functionalities to list
running services, compute average CPU/Memory usage, and track resource usage over time. See [requirements](doc/requirements.md) for more details.



## Features

- **List Instances**: Fetch from CPX concurrently and display running instances with their statuses, CPU, and memory usage.
- **Average CPU/Memory**: Calculate and display the average CPU and memory usage of services of the same type.
- **Track Resource Usage**: Continuously monitor and print CPU and memory usage of all instances of a given service at
  specified intervals.

## Installation

### Prerequisites

- Go 1.22 or later

### Steps

1. Build the executable:
    ```sh
    go build
    ```

## Usage

### Commands

- **List Instances**:
    ```sh
    ./cpx-cli instances [--healthy] [--unhealthy]
    ```
    - `--healthy`: Show only healthy instances.
    - `--unhealthy`: Show only unhealthy instances.
    - **Note**: Instances with above 80% CPU or memory usage are considered unhealthy.

- **List Services Average CPU/Memory**:
    ```sh
    ./cpx-cli services [--healthy] [--unhealthy]
    ```
    - `--healthy`: Show only healthy services.
    - `--unhealthy`: Show only unhealthy services.
    - **Note**: Services with fewer than 2
      healthy instances are considered unhealthy.

- **Track Resource Usage**:
    ```sh
    ./cpx-cli track <service_name> <interval>
    ```
    - `<service_name>`: Name of the service to track.
    - `<interval>`: Interval in seconds between each check.

### Examples

- List all instances:
    ```sh
    ./cpx-cli instances
    ```
  **Output**:
  ```
  IP          Service            Status    CPU  Memory
  10.58.1.7   RoleService        Healthy   46%  39%
  10.58.1.36  PermissionsService Healthy   41%  68%
  10.58.1.119 GeoService         Healthy   35%  22%
  10.58.1.145 AuthService        Unhealthy 82%  100%
  10.58.1.144 GeoService         Healthy   15%  63%
  10.58.1.41  RoleService        Healthy   67%  62%
  ```  
<br/>

- List only healthy instances:
    ```sh
    ./cpx-cli instances --healthy
    ```
  **Output**:
  ```
  IP          Service            Status    CPU  Memory
  10.58.1.7   RoleService        Healthy   46%  39%
  10.58.1.36  PermissionsService Healthy   41%  68%
  10.58.1.119 GeoService         Healthy   35%  22%
  10.58.1.144 GeoService         Healthy   15%  63%
  10.58.1.41  RoleService        Healthy   67%  62%
  ```
<br/>
  
- List services average CPU and memory usage:
    ```sh
    ./cpx-cli services
    ```
  **Output**:
  ```
  Service            Status    CPU Memory
  StorageService     Unhealthy 45% 51%
  IdService          Healthy   49% 46%
  TimeService        Unhealthy 40% 53%
  RoleService        Healthy   59% 49%
  PermissionsService Unhealthy 57% 43%
  GeoService         Healthy   52% 60%
  UserService        Unhealthy 52% 33%
  AuthService        Unhealthy 59% 43%
  TicketService      Unhealthy 27% 47%
  MLService          Healthy   42% 44%
  ```
<br/>

- List only healthy services average CPU and memory usage:
    ```sh
    ./cpx-cli services --healthy
    ```
  **Output**:
  ```
  Service            Status  CPU Memory
  RoleService        Healthy 44% 45%
  PermissionsService Healthy 66% 42%
  TicketService      Healthy 27% 44%
  UserService        Healthy 49% 38%
  GeoService         Healthy 56% 54%
  AuthService        Healthy 48% 44%
  StorageService     Healthy 48% 44%
  IdService          Healthy 60% 53%
  TimeService        Healthy 42% 53%
  MLService          Healthy 49% 44%
  ```
<br/>

- Track the resource usage of a specific service every 5 seconds:
    ```sh
    ./cpx-cli track PermissionsService 5
    ```
  **Output**:
  ```
  Service: PermissionsService
  Current Time: 2024-06-29 20:58:48.905373
  IP          Status    CPU Memory
  10.58.1.36  Healthy   76% 66%
  10.58.1.140 Unhealthy 18% 87%
  10.58.1.150 Healthy   35% 41%
  10.58.1.87  Healthy   44% 55%
  10.58.1.127 Healthy   59% 69%
  10.58.1.12  Healthy   52% 56%
  10.58.1.130 Unhealthy 47% 92%
  10.58.1.65  Unhealthy 86% 9%
  10.58.1.128 Healthy   27% 64%
  10.58.1.47  Healthy   48% 48%
  Service: PermissionsService
  Current Time: 2024-06-29 20:58:53.965276
  IP          Status    CPU Memory
  10.58.1.36  Healthy   66% 74%
  10.58.1.140 Healthy   25% 8%
  10.58.1.150 Healthy   42% 22%
  10.58.1.87  Unhealthy 80% 98%
  10.58.1.127 Unhealthy 88% 95%
  10.58.1.12  Unhealthy 91% 32%
  10.58.1.130 Healthy   23% 76%
  10.58.1.65  Healthy   66% 58%
  10.58.1.128 Healthy   57% 63%
  10.58.1.47  Unhealthy 91% 69%
  Service: PermissionsService
  Current Time: 2024-06-29 20:58:58.925075
  IP          Status    CPU Memory
  10.58.1.36  Healthy   77% 36%
  10.58.1.140 Healthy   34% 51%
  10.58.1.150 Healthy   16% 13%
  10.58.1.87  Healthy   57% 72%
  10.58.1.127 Healthy   1%  6%
  10.58.1.12  Unhealthy 94% 27%
  10.58.1.130 Healthy   38% 35%
  10.58.1.65  Healthy   61% 64%
  10.58.1.128 Healthy   12% 13%
  10.58.1.47  Healthy   44% 21%
  ```

## Choices and Trade-offs

- **Language**: Go was chosen for its performance and concurrency capabilities, making it suitable for real-time monitoring.
- **Libraries**: The `cobra` library is used for building the CLI due to its maturity and ease of use.

## Future Improvements

- Add more comprehensive error handling.
- Add configuration file support for custom settings.
- Enhance logging and monitoring capabilities.
