package person_controller

import "errors"

var (
    TooLongNick = errors.New("Lorem Ipsum")
    EmptyNick = errors.New("Lorem Ipsum")
    LongName = errors.New("Lorem Ipsum")
    EmptyName = errors.New("Lorem Ipsum")
    LongStackName = errors.New("Lorem Ipsum")
    InvalidBirthday = errors.New("Lorem Ipsum")
)
