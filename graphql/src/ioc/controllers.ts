import { interfaces } from 'inversify';
import Types from './types';
import {ITodosController} from "../interfaces";
import TodosController from "../api/todos/controller";

export default (c: interfaces.Container): void => {
    c.bind<ITodosController>(Types.Controllers.TodosController).to(TodosController);
};
