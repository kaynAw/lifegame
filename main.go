package main

import (
    "fmt"
    "os"
    "os/exec"
    "time"
)

type Matrix struct {
    len int
    mt  [][]bool
}

func main() {
    m := [][]bool{{
        true, true, false, false,
    }, {
        true, true, true, false,
    }, {
        false, false, true, false,
    }, {
        false, false, true, false,
    }}

    ma := Matrix{
        len: 4,
        mt:  m,
    }
    for {
        time.Sleep(time.Second/5)
        ma.expansion()
        ma.life()
        ma.printMatrix()
    }
}

func (m *Matrix)expansion() {
    need := false

    for i := 0; i < m.len;i++ {
        if m.mt[i][0] {
            need = true
            break
        }
        if m.mt[i][m.len-1] {
            need = true
            break
        }

        if m.mt[0][i] {
            need = true
            break
        }

        if m.mt[m.len-1][i] {
            need = true
            break
        }
    }

    if need {
        t := m.len + 2
        for i := 0;i < m.len;i++ {
            m.mt[i] = append(m.mt[i], false)
            m.mt[i] = append([]bool{false}, m.mt[i]...)
        }
        m.mt = append(m.mt, make([]bool, t))
        m.mt = append([][]bool{make([]bool, t)}, m.mt...)
        m.len = t
    }
}


func (m *Matrix) life() {
    cm := make([][]bool, m.len)
    for i := range cm {
        cm[i] = make([]bool, m.len)
    }

    for i := 0;i < m.len;i++ {
        for j := 0;j < m.len;j++ {
            c := 0

            if i + 1 < m.len {
                if m.mt[i+1][j] {
                    c++
                }
            }
            if i - 1 > -1 {
                if m.mt[i-1][j] {
                    c++
                }
            }

            if i + 1 < m.len && j - 1 > -1 {
                if m.mt[i+1][j-1] {
                    c++
                }
            }


            if i + 1 < m.len && j + 1 < m.len {
                if m.mt[i+1][j+1] {
                    c++
                }
            }

            if i - 1 > -1 && j - 1 > -1 {
                if m.mt[i-1][j-1] {
                    c++
                }
            }

            if i - 1 > -1 && j + 1 < m.len {
                if m.mt[i-1][j+1] {
                    c++
                }
            }

            if j - 1 > -1 {
                if m.mt[i][j-1] {
                    c++
                }
            }

            if j + 1 < m.len {
                if m.mt[i][j+1] {
                    c++
                }
            }

            if m.mt[i][j] {
                if c < 2 || c > 3 {
                    cm[i][j] = false
                } else {
                    cm[i][j] = true
                }
            } else {
                if c == 3 {
                    cm[i][j] = true
                }
            }
        }
    }
    m.mt = cm
}

func (m *Matrix) printMatrix() {
    clear()
    for i := 0; i < m.len;i++ {
        s := ""
        for j := 0;j < m.len;j++ {
            if m.mt[i][j] {
                s += `♦`
            } else {
                s += ` `
            }
        }
        fmt.Println(s)
    }
}

func clear() {
    cmd := exec.Command("clear")
    cmd.Stdout = os.Stdout
    cmd.Run()
}