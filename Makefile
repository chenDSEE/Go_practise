#=================================== config ===================================#
# path to main source code
CODE_PATH = ./cmd/bookstore


#===================================  CMD  ====================================#
debug:
	go build -gcflags="-N -l" $(CODE_PATH)

all:
	go build $(CODE_PATH)

# memoty escapes analy
memAnaly:
	go build -gcflags '-m -l' $(CODE_PATH)
