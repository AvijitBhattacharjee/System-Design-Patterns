package main

type ParkingSystem struct {
    Car map[int]int
}


func Constructor(big int, medium int, small int) ParkingSystem {
    return ParkingSystem{Car: map[int]int{1:big, 2:medium,3:small}}
}


func (this *ParkingSystem) AddCar(carType int) bool {

    if this.Car[carType] > 0 {
        this.Car[carType]-=1
        return true
    }
    return false
    
}

