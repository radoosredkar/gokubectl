1. Run kubectl get pods | grep (parameter from the commnad line)
   Example gokubectl port-forward 8123:8080 partOfPodName must:
   1. Run kubectl get pods | grep partOfPodName : DONE
   2. Get the firs pod from the returned pods (pod1) : DONE
   3. Execute kubectl port-forward  pod1 8123:8080 : DONE