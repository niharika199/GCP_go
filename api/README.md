The apis to be called

curl -X POST -d '{"Project":"proj-211305","Name":"hello","MachineType":"f1-micro","ImageProj":"rhel-cloud","Imagefamily":"rhel-7","Zone":"us-east1-b","Network":"default","Region":"us-east1"}' http://localhost:80/servercreate

curl -X POST -d '{"Name" : "hello","Project":"proj-211305","Zone":"us-east1-b"}' http://localhost:80/serverdelete

curl -X POST -d '{"Project":"proj-211305","Name":"hello"}' http://localhost:80/neworktcreate

curl -X POST -d '{"Project":"proj-211305","Name":"hello"}' http://localhost:80/networkdelete
