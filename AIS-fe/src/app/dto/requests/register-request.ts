export class RegisterRequest {
    firstName: string;
    lastName: string;
    password: string;
//  TODO: Add more attributes if needed

    constructor(firstName: string, lastName: string, password: string) {
        this.firstName = firstName;
        this.lastName = lastName;
        this.password = password;
    }
}