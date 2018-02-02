import { interfaces } from 'inversify';
import Types from '../ioc/types';
import {ITodoRepository} from "../interfaces";
import TodoRepository from "../api/todos/repository";

export default (c: interfaces.Container): void => {
    c.bind<ITodoRepository>(Types.Repositories.TodoRepository).to(TodoRepository);
};
