package driver

/*
#cgo LDFLAGS: -lcomedi -lm
#include "io.h"
#include "channels.h"
#include "elev.h"
*/
import "C"

type elev_button_type_t int

type elev_motor_direction_t int

const (
    BUTTON_CALL_UP elev_button_type_t = iota
    BUTTON_CALL_DOWN
    BUTTON_COMMAND
)

const (
    DIRN_DOWN elev_motor_direction_t = -1
    DIRN_STOP = 0
    DIRN_UP = 1
)

func Init() int {
    return int(C.elev_init())
    go updateFloorLights()
}

func setMotorDir(dir elev_motor_direction_t) { // made private
    C.elev_set_motor_direction(C.elev_motor_direction_t(dir))
}

func GetFloorSensor() int{
    return int(C.elev_get_floor_sensor_signal())
}

func GetButtonSignal(button elev_button_type_t, floor int) int {
    return int(C.elev_get_button_signal(C.elev_button_type_t(button), C.int(floor)))
}

func GetStopSignal() int {
    return int(C.elev_get_stop_signal())
}

func GetObstructionSignal() int {
    return int(C.elev_get_stop_signal())
}

func setFloorIndicator(floor int){ // made private
    C.elev_set_floor_indicator(C.int(floor))
}

func SetButtonLamp(button elev_button_type_t, floor int, value int){
    C.elev_set_button_lamp(C.elev_button_type_t(button), C.int(floor), C.int(value))
}

func SetStopLamp(value int){
    C.elev_set_stop_lamp(C.int(value))
}

func SetDoorOpenLamp(value int){
    C.elev_set_door_open_lamp(C.int(value));
}

func GoToFloor(desiredFloor int) {
    currentFloor := GetFloorSensor()
    if desiredFloor == currentFloor {
        return
        } else {
            setMotorDir(desiredFloor - currentFloor)
            for desiredFloor != currentFloor {
                // Weit for how long?
        }
        return
    }
}

func updateFloorLights(){ // Run as go routine from Init()
    currentFloor := GetFloorSensor()
    // if currentFloor
    // setFloorIndicator(currentFloor)
}



// Legg til i init: Kjør ned til et nivå    



/*
Public functions:
GoToFloor(floor int)
GetFloorSensor
ButtonPushedOnFloor() chan


buttonpress channel:

const (
floorButton3down type_t = iota
floorButton2up
floorButton2down
floorButton1up
floorButton1down
floorButton0up
elevatorButtonCommand3
elevatorButtonCommand2
elevatorButtonCommand1
elevatorButtonCommand0
elevatorButtonStop

go routine som sjekker for knapper "hele tiden" og putter dem ut på en channel?
*/




