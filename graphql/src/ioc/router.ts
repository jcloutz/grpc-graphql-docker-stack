import { interfaces } from 'inversify';
import { IServerRegisterable } from '../interfaces';
import Types from '../ioc/types';
import TodoRoutes from "../api/todos/routes";

export default (c: interfaces.Container): void => {
    c.bind<IServerRegisterable>(Types.Routes.Route).to(TodoRoutes);
};
