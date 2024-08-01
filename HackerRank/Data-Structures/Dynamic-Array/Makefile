CXX=g++
CXXFLAGS=-std=c++11 -ggdb -Wno-deprecated-declarations
SRCS=solution.cpp
BINS=dynamic_array.bin

all: $(BINS)

$(BINS): $(SRCS)
	$(CXX) $(CXXFLAGS) -o $@ $^

clean:
	rm -f $(BINS)
