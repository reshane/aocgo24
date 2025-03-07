package utils

import (
    "os"
    "log"
    "bufio"
)

type io struct{}
var Io io

func (io) ReadToString(path string) (string, error) {
    content, err := os.ReadFile(path)
    if err != nil {
        log.Printf("Error: Could not read file %s: %s", path, err)
        return "", err
    }
    return string(content), nil
}

func (io) ReadToLines(path string) ([]string, error) {
    data := make([]string, 0, 1000)

    file, err := os.Open(path)
    if err != nil {
        log.Printf("Error: Could not open file %s: %s", path, err)
        return data, err
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        data = append(data, scanner.Text())
    }

    if err = scanner.Err(); err != nil {
        log.Printf("Error: could not read from file %s: %s", path, err)
    }

    return data, nil
}

func (io) ReadToRunes(path string) ([]rune, error) {
    data := make([]byte, 0, 1000)

    file, err := os.Open(path)
    if err != nil {
        log.Printf("Error: Could not open file %s: %s", path, err)
        return []rune(string(data)), err
    }
    defer file.Close()

    reader := bufio.NewReader(file)
    for {
        char, err := reader.ReadByte()
        if err != nil {
            break
        }
        data = append(data, char)
    }

    return []rune(string(data)), nil
}
