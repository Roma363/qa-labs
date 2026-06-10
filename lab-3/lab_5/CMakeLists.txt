cmake_minimum_required(VERSION 3.20)

project(lab_5 LANGUAGES CXX)

set(CMAKE_CXX_STANDARD 17)
set(CMAKE_CXX_STANDARD_REQUIRED ON)
set(CMAKE_CXX_EXTENSIONS OFF)

add_executable(lab_5_server
        src/server.cpp
        src/cache.cpp
)

target_include_directories(lab_5_server PRIVATE
        ${CMAKE_CURRENT_SOURCE_DIR}/include
)

target_link_libraries(lab_5_server PRIVATE
        ws2_32
)

add_custom_command(TARGET lab_5_server POST_BUILD
        COMMAND ${CMAKE_COMMAND} -E copy_directory
        ${CMAKE_CURRENT_SOURCE_DIR}/public
        $<TARGET_FILE_DIR:lab_5_server>/public
)