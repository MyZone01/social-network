export class Login {
    email: string;
    password: string;

    constructor(data: any) {
        this.email = data.email;
        this.password = data.password;
    }

    static isEmail(email: string): boolean {
        const re = /\S+@\S+\.\S+/;
        return re.test(email);
    }

    validate(): [boolean, string] {
        if (!Login.isEmail(this.email)) {
            return [false, "Invalid email"];
        }

        if (!this.password) {
            return [false, "Password is missing. Please provide it."];
        }

        return [true, "User data is valid"];
    }
}