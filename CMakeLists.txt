cmake_minimum_required(VERSION 3.10)
project(BlockDecoder)

set(CMAKE_CXX_STANDARD 14)  # Или более позднюю версию

add_executable(BlockDecoder
        main.cpp

        src/coder.cpp
        src/channel.cpp
        src/orchestrator.cpp
        src/gen_data.cpp

        src/coder.h
        src/channel.h
        src/orchestrator.h
        src/gen_data.h)