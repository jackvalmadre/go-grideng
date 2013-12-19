/*
Distributed map-reduce on a Grid Engine system.

To avoid manipulating code, the paradigm of the package is to create an executable which invokes itself with different command-line flags.

Defining Tasks

Grid Engine tasks can be defined by implementing the Task interface directly, but usually you would use Func or ReduceFunc to wrap a function.

Examples

An example of how to use the package for a simple map operation:
	func main() {
		grideng.Register("square", grideng.Func(func(x float64) float64 { return x * x }))
		flag.Parse()
		grideng.ExecIfSlave()

		const n = 100
		x := make([]float64, n)
		for i := range x {
			x[i] = float64(i + 1)
		}

		y := make([]float64, len(x))
		if err := grideng.Map("square", y, x, nil); err != nil {
			fmt.Fprintln(os.Stderr, "map:", err)
		}
	}

Note that this adds several command line flags:
	$ ./example -help
	Usage of ./example:
	  -square.l="": Resource flag (-l) to qsub
	  -grideng.addr="": Address of master on network.
	  -grideng.task="": Task to execute as slave. Empty to execute as master.
The address is used in both net.Listen() in the master and net.Dial() in the slaves.
The <task>.l flag can be used to configure qsub resources e.g. h_vmem and virtual_free.

To call a function which accepts a constant parameter for all x[i]:
	grideng.Register("pow", grideng.Func(math.Pow))
	// ...
	y := make([]float64, len(x))
	err := grideng.Map("pow", y, x, float64(2))

To do a reduce operation:
	grideng.Register("add", grideng.ReduceFunc(func(x, y float64) float64 { return x + y }))
	// ...
	var total float64
	err := grideng.Reduce("add", &total, x, nil)
*/
package grideng