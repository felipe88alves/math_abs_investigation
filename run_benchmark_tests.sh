#! /bin/bash

set -eo pipefail

# ------------- Start Functions -------------

function create_dir () {
    i=1
    while [ -d "$results_dir" ]; do
        results_dir=${results_dir%_*}_${i}
        i=$((i+1))
    done
    mkdir ${results_dir}
}

function run_tests_per_var_type () {
    var_type=$1
    local output_dir=${results_dir}/${var_type}
    mkdir ${output_dir}
    printf "      Running Benchmarking Tests for ${var_type} - Current Abs Function\n"
    gotip test -bench ^BenchmarkAbs_${var_type}$ -cpu 1 -benchmem -benchtime ${num_benchmark_executions}x -count ${num_test_count} > ${output_dir}/BenchmarkAbs.txt
    printf "      Running Benchmarking Tests for ${var_type} - Abs Function with Generic Parameters\n"
    gotip test -bench ^BenchmarkGenericAbsReturnFloat_${var_type}$ -cpu 1 -benchmem -benchtime ${num_benchmark_executions}x -count ${num_test_count} > ${output_dir}/BenchmarkGenericAbsReturnFloat.txt
    printf "      Running Benchmarking Tests for ${var_type} - Abs Function with Generic Parameters and return\n"
    gotip test -bench ^BenchmarkGenericAbsReturnT_${var_type}$ -cpu 1 -benchmem -benchtime ${num_benchmark_executions}x -count ${num_test_count} > ${output_dir}/BenchmarkGenericAbsReturnT.txt
    printf "      Compiling Benchmarking Tests Results for ${var_type}\n"
    benchstat ${output_dir}/BenchmarkAbs.txt ${output_dir}/BenchmarkGenericAbsReturnFloat.txt ${output_dir}/BenchmarkGenericAbsReturnT.txt > ${results_dir}/${var_type}Benchstat.txt
}
# ------------- End Functions -------------

# ------------- Start main script -------------

  # Define variables
  declare results_dir="benchmark_results_1"
  declare num_test_count=100
  declare num_benchmark_executions=1000000000
  
  # init
  create_dir
  printf "Running Benchmarking Tests for float64\n"
  run_tests_per_var_type float64
  printf "Running Benchmarking Tests for int\n"
  run_tests_per_var_type int
  printf "Running Benchmarking Tests for int64\n"
  run_tests_per_var_type int64
  printf "Running Benchmarking Tests for Struct float64\n"
  run_tests_per_var_type structFloat64
  printf "Running Benchmarking Tests for Struct int64\n"
  run_tests_per_var_type structInt64
  printf "Running Benchmarking Tests for Custom Type float64\n"
  run_tests_per_var_type typeFloat64
  printf "Running Benchmarking Tests for Custom Type int64\n"
  run_tests_per_var_type typeInt64
  printf "Test results output to ${results_dir}\n"s
