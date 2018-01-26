import { IUser } from '../interfaces';

export default class User implements IUser {
    public _id: string;
    public age: number;
    public name: string;
    public lastName: string;
}
