CXX=g++
CXXFLAGS=-std=c++11 -Wno-deprecated-declarations
SRCS=solution.cpp
BINS=reverse_array.bin

all: $(BINS)

$(BINS): $(SRCS)
	$(CXX) $(CXXFLAGS) -o $@ $^

clean:
	rm -f $(BINS)
