export interface IUser {
    username: string;
    role: string;
    token: string;
    picture?: string;
}

export interface IUserFormValues {
    email: string;
    password: string;
    displayName?: string;
    username?: string;
}
