import { inject, injectable } from 'inversify';
import BaseRouter from 'app/common/base-router';
import { IRouteConfiguration, ITodosController } from 'app/interfaces';
import Types from 'app/ioc/types';
import validate from 'app/api/todos/validate';
import TodosController from "app/api/todos/controller";
import * as Joi from "joi";

@injectable()
export default class TodosRoutes extends BaseRouter {
    private controller: ITodosController;

    constructor(@inject(Types.Controllers.TodosController) controller: TodosController) {
        super();

        this.controller = controller;
    }

    public routes(): IRouteConfiguration[] {
        return [
            {
                method: 'POST',
                path: '/api/todos',
                config: {
                    handler: this.controller.create,
                    validate: validate.create,
                    description: 'Method that creates a new todo.',
                    tags: ['api', 'todos'],
                    auth: false,
                    response: {
                        status: {
                            201: Joi.object({
                                status: 201,
                                data: Joi.object({
                                    id: Joi.string().required(),
                                    name: Joi.string().required(),
                                    completed: Joi.boolean().required(),
                                    createdAt: Joi.date().required(),
                                    updatedAt: Joi.date().required(),
                                }),
                            }),
                            400: Joi.object({
                                statusCode: 400,
                                error: Joi.string().required(),
                                message: Joi.string().required(),
                                validation: Joi.object({
                                    source: Joi.string().required(),
                                    keys: Joi.array().items(Joi.string()),
                                }),
                            }),
                        },
                    },
                },
            },
            // {
            //     method: 'PUT',
            //     path: '/api/todos/{id}',
            //     config: {
            //         handler: this.controller.updateById,
            //         validate: validate.updateById,
            //         description: 'Method that updates a todo by its id.',
            //         tags: ['api', 'todos'],
            //         auth: false,
            //     },
            // },
            {
                method: 'GET',
                path: '/api/todos/{id}',
                config: {
                    handler: this.controller.getById,
                    validate: validate.getById,
                    description: 'Method that get a todo by its id.',
                    tags: ['api', 'todos'],
                    auth: false,
                },
            },
            {
                method: 'GET',
                path: '/api/todos',
                config: {
                    handler: this.controller.getAll,
                    description: 'Method that gets all todos.',
                    tags: ['api', 'todos'],
                    auth: false,
                },
            },
            // {
            //     method: 'DELETE',
            //     path: '/api/todos/{id}',
            //     config: {
            //         handler: this.controller.deleteById,
            //         validate: validate.deleteById,
            //         description: 'Method that deletes a todo by its id.',
            //         tags: ['api', 'todos'],
            //         auth: false,
            //     },
            // },
        ];
    }
}
