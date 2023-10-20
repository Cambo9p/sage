# sage
Speech Activated Guide and Expert

@startuml
namespace python {
    class speech_to_text {
        +read_from_microphone()
        -write_data(data)
    }
    class client {
        
    } 
}

namespace go {

class server {
}

class filter {
}

class notifier {
}


}

client -- server : TCP
server --> filter
filter --> notifier
@enduml
