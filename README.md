<p align="right">
 <img src="https://github.com/hellgate75/go-invite-customers/workflows/Go/badge.svg?branch=master"></img>
&nbsp;&nbsp;<img src="https://api.travis-ci.com/hellgate75/go-invite-customers.svg?branch=master" alt="trevis-ci" width="98" height="20" />&nbsp;&nbsp;<a href="https://travis-ci.com/hellgate75/go-invite-customers">Check last build on Travis-CI</a>
 </p>

# go-invite-customers
Defines a command that reads customer data from a file or a pipe stream and collects all user data in a specific distance range from home coordinates


## Command line options

Command syntax is following:
```
go-invite-customers -[param0]=value0 ...  -[paramN]=valueN
Parameters:
  -detailed
        Create Output for invited and excluded, instead of only invited customers
  -distance float
        Max distance from base coordinate (default 100)
  -in-enc string
        Input encoding format: [json yaml xml] (default "json")
  -input string
        Given file, url or pipe that contains data
  -latitude float
        Base latitude degrees in float number [W is negative] (default 53.339428)
  -longitude float
        Base longitude degrees in float number [S is negative] (default -6.257664)
  -out-enc string
        Output encoding format: [text json yaml xml] (default "text")
  -per-line-input
        Use one read line in input for parsing the data, instead of reading the list (default true)
  -silent
        Execute silent output
  -unit string
        Measure Unit for distance [K is for Kilometers, M is for Miles and N is for Nautical Miles] (default "K")
```

### Agument details

Following arguments can be passed to the command line:

* `[-detailed]` - If true print in output invited and excluded users, or if false only invited users
* `[-distance]` - Specify maximum distance for customer office from the base coordinates
* `[-unit]` - Specify the measure unit for the distance (K: Kms, M: Mls, N, NMls)
* `[-silent]` - Execute a silent execution
* `[-latitude]` - Base office latitude in degrees, with positive (E) or negative (W) values
* `[-longitude]` - Base office logitude in degrees, with positive (N) or negative (S) values
* `[-per-line-input]` - Define the kind of imput from the stream (see input data type samples)
* `[-in-enc]` - Input stream encoding format
* `[-out-enc]` - Output text encoding format
* `[-input]` - Defines the imput stream : udp://host:port, tcp:host:port, [http, https]://host[:port]/.. or any other format is considered as a file path


### Input Data Types Samples

Data types can be collected by :

Line (each line contains data formats), and the code declaration is following one (model.CustomerOffice):

```
type CustomerOffice struct {
	UserId    int64  `json:"user_id,omitempty" yaml:"user_id,omitempty" xml:"user-id,omitempty"`
	Name      string `json:"name,omitempty" yaml:"name,omitempty" xml:"name,omitempty"`
	Latitude  string `json:"latitude,omitempty" yaml:"latitude,omitempty" xml:"latitude,omitempty"`
	Longitude string `json:"longitude,omitempty" yaml:"longitude,omitempty" xml:"longitude,omitempty"`
}
```

Single element (all stream returns a single emelement) and the declaration is following one (model.CustomerOfficeList):

```
type CustomerOfficeList struct {
	List []CustomerOffice `json:"customers,omitempty" yaml:"customers,omitempty" xml:"customers,omitempty"`
}

type CustomerOffice struct {
	UserId    int64  `json:"user_id,omitempty" yaml:"user_id,omitempty" xml:"user-id,omitempty"`
	Name      string `json:"name,omitempty" yaml:"name,omitempty" xml:"name,omitempty"`
	Latitude  string `json:"latitude,omitempty" yaml:"latitude,omitempty" xml:"latitude,omitempty"`
	Longitude string `json:"longitude,omitempty" yaml:"longitude,omitempty" xml:"longitude,omitempty"`
}
```



## Opertions

Following commands available after installation and configuration of [Google Go](https://golang.org/doc/install)


### Get the command

Command is installed in `$GOPATH`\bin folder running:

```
go get -u github.com/hellgate75/go-invite-customers
```



### Test the command

Test is executed running:

```
go test -cover github.com/hellgate75/go-invite-customers
```


### Rebuild the command

Test is executed running:

```
go build -buildmode=exe github.com/hellgate75/go-invite-customers
```




## License

The library is licensed with [LGPL v. 3.0](/LICENSE) clauses, with prior authorization of author before any production or commercial use. Use of this library or any extension is prohibited due to high risk of damages due to improper use. No warranty is provided for improper or unauthorized use of this library or any implementation.

Any request can be prompted to the author [Fabrizio Torelli](https://www.linkedin.com/in/fabriziotorelli) at the following email address:

[hellgate75@gmail.com](mailto:hellgate75@gmail.com)
 
