#!/bin/bash
# Athera Language Runner - Easy launcher for Athera programs

echo "╔══════════════════════════════════════════╗"
echo "║   Athera Programming Language Runner    ║"
echo "╚══════════════════════════════════════════╝"
echo ""

if [ $# -eq 0 ]; then
    echo "Available example programs:"
    echo ""
    echo "BASIC EXAMPLES:"
    echo "  1. morning_routine.ath - Original example"
    echo "  2. complete_example.ath - Full feature showcase"
    echo "  3. demo.ath - Visual demonstration"
    echo ""
    echo "ADVANCED EXAMPLES:"
    echo "  4. advanced_features.ath - Parameters, errors, parallel"
    echo "  5. module_demo.ath - Custom modules"
    echo "  6. error_handling_demo.ath - Error handling patterns"
    echo "  7. parallel_demo.ath - Concurrent execution"
    echo "  8. real_world_app.ath - Complete application"
    echo ""
    echo "FILE OPERATIONS:"
    echo "  9. real_backup_test.ath - Real file operations"
    echo "  10. file_test.ath - File handling demo"
    echo ""
    echo "ALGORITHMS:"
    echo "  11. fibonacci.ath - Algorithm example"
    echo ""
    echo "Usage: ./run_athera.sh <filename.ath>"
    echo "   Or: ./run_athera.sh <number>"
    echo ""
    echo "Examples:"
    echo "  ./run_athera.sh morning_routine.ath"
    echo "  ./run_athera.sh 4"
    exit 0
fi

# Handle numeric shortcuts
case "$1" in
    1) FILE="morning_routine.ath" ;;
    2) FILE="complete_example.ath" ;;
    3) FILE="demo.ath" ;;
    4) FILE="advanced_features.ath" ;;
    5) FILE="module_demo.ath" ;;
    6) FILE="error_handling_demo.ath" ;;
    7) FILE="parallel_demo.ath" ;;
    8) FILE="real_world_app.ath" ;;
    9) FILE="real_backup_test.ath" ;;
    10) FILE="file_test.ath" ;;
    11) FILE="fibonacci.ath" ;;
    *) FILE="$1" ;;
esac

echo "Running: $FILE"
echo "═══════════════════════════════════════════"
echo ""

python3 athera_interpreter.py "$FILE"

echo ""
echo "═══════════════════════════════════════════"
echo "Athera execution complete!"
