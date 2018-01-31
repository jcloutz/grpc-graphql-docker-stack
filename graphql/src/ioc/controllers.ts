import { interfaces } from 'inversify';
import ArticlesController from 'app/api/articles/controller';
import UsersController from 'app/api/users/controller';
import { IArticlesController, IUsersController } from 'app/interfaces';
import Types from './types';
import {ITodosController} from "../interfaces";
import TodosController from "../api/todos/controller";

export default (c: interfaces.Container): void => {
    c.bind<IUsersController>(Types.Controllers.UsersController).to(UsersController);
    c.bind<IArticlesController>(Types.Controllers.ArticlesController).to(ArticlesController);
    c.bind<ITodosController>(Types.Controllers.TodosController).to(TodosController);
};
