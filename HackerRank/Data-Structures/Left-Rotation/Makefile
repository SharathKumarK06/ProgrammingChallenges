CXX=g++
CXXFLAGS=-std=c++11 -Wno-deprecated-declarations
SRCS=solution.cpp
BINS=left_rotation.bin

all: $(BINS)

$(BINS): $(SRCS)
	$(CXX) $(CXXFLAGS) -o $@ $^

clean:
	rm -f $(BINS)
