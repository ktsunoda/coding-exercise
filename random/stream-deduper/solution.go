// http://repl.it/rx4/39
// Stream deduper
package main

import "errors"
import "fmt"

type Stream struct {
    data []int
}

func NewStream() *Stream {
    return &Stream{data:[]int{1, 1, 2, 3, 3, 3, 4, 4, 4, 4, 5, 6, 7, 7, 7, 7, 8}}
    // return &Stream{data:[]int{1, 1, 2, 3, 3, 3}}
}

func (s *Stream) HasNext() bool {
    return len(s.data) > 0
}

func (s *Stream) GetNext() int {
    next := s.data[0]
    s.data = s.data[1:]
    
    return next
}

func (s *Stream) AppendData(data []int) {
    s.data = append(s.data, data...)
}

type MyStream struct {
    stream *Stream
    
    previous *int
    next *int
}

func NewMyStream() *MyStream {
    return &MyStream{stream:NewStream()}
}

func (m *MyStream) HasNext() bool {
    if m.next == nil && m.stream.HasNext() {
        for m.stream.HasNext() {
            next := m.stream.GetNext()
        
            if m.previous == nil || *m.previous != next {
                m.next = new(int)
                *m.next = next
                break
            }
        }
    }
    
    return m.next != nil
}

func (m *MyStream) GetNext() (int, error) {
    var next int
    err := errors.New("Success")
    
    if m.next != nil {
        next = *m.next
        m.previous = new(int)
        *m.previous = next
        m.next = nil
    } else {
        err = errors.New("No data")
    }

    return next, err
}

func (m *MyStream) AppendData(data []int) {
    m.stream.AppendData(data)
}

func main() {
    myStream := NewMyStream()
    
    for myStream.HasNext() {
        fmt.Println(myStream.GetNext())
    }
    
    fmt.Println(myStream.GetNext())

    myStream.AppendData([]int{8, 9, 9, 10, 11, 12, 12})
    
    for myStream.HasNext() {
        fmt.Println(myStream.GetNext())
    }
    
    fmt.Println(myStream.GetNext())
}
