import { inject, injectable } from 'inversify';
import BaseRouter from 'app/common/base-router';
import { IRouteConfiguration, IUsersController } from 'app/interfaces';
import Types from 'app/ioc/types';
import validate from 'app/api/users/validate';

@injectable()
export default class UserRoutes extends BaseRouter {
    private controller: IUsersController;

    constructor(@inject(Types.Controllers.UsersController) controller: IUsersController) {
        super();

        this.controller = controller;
    }

    public routes(): IRouteConfiguration[] {
        return [
            {
                method: 'POST',
                path: '/api/users',
                config: {
                    handler: this.controller.create,
                    validate: validate.create,
                    description: 'Method that creates a new user.',
                    tags: ['api', 'users'],
                    auth: false,
                },
            },
            {
                method: 'PUT',
                path: '/api/users/{id}',
                config: {
                    handler: this.controller.updateById,
                    validate: validate.updateById,
                    description: 'Method that updates a user by its id.',
                    tags: ['api', 'users'],
                    auth: false,
                },
            },
            {
                method: 'GET',
                path: '/api/users/{id}',
                config: {
                    handler: this.controller.getById,
                    validate: validate.getById,
                    description: 'Method that get a user by its id.',
                    tags: ['api', 'users'],
                    auth: false,
                },
            },
            {
                method: 'GET',
                path: '/api/users',
                config: {
                    handler: this.controller.getAll,
                    description: 'Method that gets all users.',
                    tags: ['api', 'users'],
                    auth: false,
                },
            },
            {
                method: 'DELETE',
                path: '/api/users/{id}',
                config: {
                    handler: this.controller.deleteById,
                    validate: validate.deleteById,
                    description: 'Method that deletes a user by its id.',
                    tags: ['api', 'users'],
                    auth: false,
                },
            },
        ];
    }
}