# Accelerator Log

## Options
```json
{
  "projectName" : "go-service-http-rabbit-dev-001"
}
```
## Log
```
┏ engine (Chain)
┃  Info Running Chain(Combo, UniquePath)
┃ ┏ engine.transformations[0] (Combo)
┃ ┃  Info Combo running as Chain(Merge(merge), UniquePath(UseLast))
┃ ┃ engine.transformations[0].merge (Chain)
┃ ┃  Info Running Chain(Merge, UniquePath)
┃ ┃ ┏ engine.transformations[0].merge.transformations[0] (Merge)
┃ ┃ ┃  Info Running Merge(Combo, Combo)
┃ ┃ ┃ ┏ engine.transformations[0].merge.transformations[0].sources[0] (Combo)
┃ ┃ ┃ ┃  Info Combo running as Chain(Include, Exclude)
┃ ┃ ┃ ┃ engine.transformations[0].merge.transformations[0].sources[0].<combo> (Chain)
┃ ┃ ┃ ┃  Info Running Chain(Include, Exclude)
┃ ┃ ┃ ┃ ┏ engine.transformations[0].merge.transformations[0].sources[0].<combo>.transformations[0] (Include)
┃ ┃ ┃ ┃ ┃  Info Will include [**/*]
┃ ┃ ┃ ┃ ┃ Debug .gitignore matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ Debug README.md matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ Debug go.mod matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ Debug go.sum matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ Debug icons/go.png matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ Debug main.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/go-chi/chi/.gitignore matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/go-chi/chi/CHANGELOG.md matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/go-chi/chi/CONTRIBUTING.md matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/go-chi/chi/LICENSE matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/go-chi/chi/Makefile matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/go-chi/chi/README.md matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/go-chi/chi/chain.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/go-chi/chi/chi.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/go-chi/chi/context.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/go-chi/chi/mux.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/go-chi/chi/tree.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/nebhale/client-go/LICENSE matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/nebhale/client-go/NOTICE matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/nebhale/client-go/bindings/binding.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/nebhale/client-go/bindings/bindings.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/nebhale/client-go/internal/secret.go matched [**/*] -> included
┃ ┃ ┃ ┃ ┃ Debug vendor/modules.txt matched [**/*] -> included
┃ ┃ ┃ ┃ ┗ Debug workload.yaml matched [**/*] -> included
┃ ┃ ┃ ┃ ┏ engine.transformations[0].merge.transformations[0].sources[0].<combo>.transformations[1] (Exclude)
┃ ┃ ┃ ┃ ┃  Info Will exclude [icons, workload.yaml, accelerator.yaml, Tiltfile, go.mod]
┃ ┃ ┃ ┃ ┃ Debug .gitignore didn't match [icons, workload.yaml, accelerator.yaml, Tiltfile, go.mod] -> included
┃ ┃ ┃ ┃ ┃ Debug README.md didn't match [icons, workload.yaml, accelerator.yaml, Tiltfile, go.mod] -> included
┃ ┃ ┃ ┃ ┃ Debug go.mod matched [icons, workload.yaml, accelerator.yaml, Tiltfile, go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ Debug go.sum didn't match [icons, workload.yaml, accelerator.yaml, Tiltfile, go.mod] -> included
┃ ┃ ┃ ┃ ┃ Debug icons/go.png didn't match [icons, workload.yaml, accelerator.yaml, Tiltfile, go.mod] -> included
┃ ┃ ┃ ┃ ┃ Debug main.go didn't match [icons, workload.yaml, accelerator.yaml, Tiltfile, go.mod] -> included
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/go-chi/chi/.gitignore didn't match [icons, workload.yaml, accelerator.yaml, Tiltfile, go.mod] -> included
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/go-chi/chi/CHANGELOG.md didn't match [icons, workload.yaml, accelerator.yaml, Tiltfile, go.mod] -> included
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/go-chi/chi/CONTRIBUTING.md didn't match [icons, workload.yaml, accelerator.yaml, Tiltfile, go.mod] -> included
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/go-chi/chi/LICENSE didn't match [icons, workload.yaml, accelerator.yaml, Tiltfile, go.mod] -> included
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/go-chi/chi/Makefile didn't match [icons, workload.yaml, accelerator.yaml, Tiltfile, go.mod] -> included
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/go-chi/chi/README.md didn't match [icons, workload.yaml, accelerator.yaml, Tiltfile, go.mod] -> included
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/go-chi/chi/chain.go didn't match [icons, workload.yaml, accelerator.yaml, Tiltfile, go.mod] -> included
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/go-chi/chi/chi.go didn't match [icons, workload.yaml, accelerator.yaml, Tiltfile, go.mod] -> included
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/go-chi/chi/context.go didn't match [icons, workload.yaml, accelerator.yaml, Tiltfile, go.mod] -> included
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/go-chi/chi/mux.go didn't match [icons, workload.yaml, accelerator.yaml, Tiltfile, go.mod] -> included
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/go-chi/chi/tree.go didn't match [icons, workload.yaml, accelerator.yaml, Tiltfile, go.mod] -> included
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/nebhale/client-go/LICENSE didn't match [icons, workload.yaml, accelerator.yaml, Tiltfile, go.mod] -> included
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/nebhale/client-go/NOTICE didn't match [icons, workload.yaml, accelerator.yaml, Tiltfile, go.mod] -> included
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/nebhale/client-go/bindings/binding.go didn't match [icons, workload.yaml, accelerator.yaml, Tiltfile, go.mod] -> included
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/nebhale/client-go/bindings/bindings.go didn't match [icons, workload.yaml, accelerator.yaml, Tiltfile, go.mod] -> included
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/nebhale/client-go/internal/secret.go didn't match [icons, workload.yaml, accelerator.yaml, Tiltfile, go.mod] -> included
┃ ┃ ┃ ┃ ┃ Debug vendor/modules.txt didn't match [icons, workload.yaml, accelerator.yaml, Tiltfile, go.mod] -> included
┃ ┃ ┃ ┗ ┗ Debug workload.yaml matched [icons, workload.yaml, accelerator.yaml, Tiltfile, go.mod] -> excluded
┃ ┃ ┃ ┏ engine.transformations[0].merge.transformations[0].sources[1] (Combo)
┃ ┃ ┃ ┃  Info Combo running as Chain(Include, Chain(chain))
┃ ┃ ┃ ┃ engine.transformations[0].merge.transformations[0].sources[1].<combo> (Chain)
┃ ┃ ┃ ┃  Info Running Chain(Include, Chain)
┃ ┃ ┃ ┃ ┏ engine.transformations[0].merge.transformations[0].sources[1].<combo>.transformations[0] (Include)
┃ ┃ ┃ ┃ ┃  Info Will include [workload.yaml, Tiltfile, go.mod]
┃ ┃ ┃ ┃ ┃ Debug .gitignore didn't match [workload.yaml, Tiltfile, go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ Debug README.md didn't match [workload.yaml, Tiltfile, go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ Debug go.mod matched [workload.yaml, Tiltfile, go.mod] -> included
┃ ┃ ┃ ┃ ┃ Debug go.sum didn't match [workload.yaml, Tiltfile, go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ Debug icons/go.png didn't match [workload.yaml, Tiltfile, go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ Debug main.go didn't match [workload.yaml, Tiltfile, go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/go-chi/chi/.gitignore didn't match [workload.yaml, Tiltfile, go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/go-chi/chi/CHANGELOG.md didn't match [workload.yaml, Tiltfile, go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/go-chi/chi/CONTRIBUTING.md didn't match [workload.yaml, Tiltfile, go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/go-chi/chi/LICENSE didn't match [workload.yaml, Tiltfile, go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/go-chi/chi/Makefile didn't match [workload.yaml, Tiltfile, go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/go-chi/chi/README.md didn't match [workload.yaml, Tiltfile, go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/go-chi/chi/chain.go didn't match [workload.yaml, Tiltfile, go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/go-chi/chi/chi.go didn't match [workload.yaml, Tiltfile, go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/go-chi/chi/context.go didn't match [workload.yaml, Tiltfile, go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/go-chi/chi/mux.go didn't match [workload.yaml, Tiltfile, go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/go-chi/chi/tree.go didn't match [workload.yaml, Tiltfile, go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/nebhale/client-go/LICENSE didn't match [workload.yaml, Tiltfile, go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/nebhale/client-go/NOTICE didn't match [workload.yaml, Tiltfile, go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/nebhale/client-go/bindings/binding.go didn't match [workload.yaml, Tiltfile, go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/nebhale/client-go/bindings/bindings.go didn't match [workload.yaml, Tiltfile, go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/github.com/nebhale/client-go/internal/secret.go didn't match [workload.yaml, Tiltfile, go.mod] -> excluded
┃ ┃ ┃ ┃ ┃ Debug vendor/modules.txt didn't match [workload.yaml, Tiltfile, go.mod] -> excluded
┃ ┃ ┃ ┃ ┗ Debug workload.yaml matched [workload.yaml, Tiltfile, go.mod] -> included
┃ ┃ ┃ ┃ ┏ engine.transformations[0].merge.transformations[0].sources[1].<combo>.transformations[1] (Chain)
┃ ┃ ┃ ┃ ┃  Info Running Chain(ReplaceText)
┃ ┃ ┃ ┃ ┃ ┏ engine.transformations[0].merge.transformations[0].sources[1].<combo>.transformations[1].transformations[0] (ReplaceText)
┃ ┃ ┗ ┗ ┗ ┗  Info Will replace [go-service-http-rabbit->go-service-http-rabb...(truncated)]
┃ ┗ ╺ engine.transformations[0].merge.transformations[1] (UniquePath)
┗ ╺ engine.transformations[1] (UniquePath)
```
