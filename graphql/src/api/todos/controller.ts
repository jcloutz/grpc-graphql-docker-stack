import { inject, injectable } from 'inversify';
import { IArticle, IArticleResolver, IUsersController } from 'app/interfaces';
import Types from 'app/ioc/types';
import {ILogger, ITodo, ITodoEdit, ITodoRepository, ITodosController} from "../../interfaces";
import { Request, ReplyNoContinue } from "hapi";
import * as Boom from 'boom';

@injectable()
export default class TodosController implements ITodosController {
    // public create: (request: Request, response: ReplyNoContinue) => Promise<any>;
    public updateById: (request: Request, response: ReplyNoContinue) => Promise<any>;
    // public getByid: (request: request, response: replynocontinue) => promise<any>;
    // public getAll: (request: Request, response: ReplyNoContinue) => Promise<any>;
    public deleteById: (request: Request, response: ReplyNoContinue) => Promise<any>;
    public bulkUpdate: (request: Request, response: ReplyNoContinue) => Promise<any>;
    public bulkDelete: (request: Request, response: ReplyNoContinue) => Promise<any>;

    protected repository: ITodoRepository;

    constructor(
        @inject(Types.Repositories.TodoRepository) repository: ITodoRepository,
    ) {
        this.repository = repository;
    }

    public getAll = async (request: Request, response: ReplyNoContinue): Promise<any> => {
        try {
            const todos = await this.repository.getAll();

            return response({
                status: 200,
                data: todos,
            });
        } catch (error) {
            return response(Boom.badImplementation(error));
        }
    };

    public getById = async (request: Request, response: ReplyNoContinue): Promise<any> => {
        try {
            const id = decodeURIComponent(request.params.id);

            const todo = await this.repository.getById(id);

            return response({
                status: 200,
                data: todo,
            });
        } catch (error) {
            return response(Boom.badImplementation(error));
        }
    }

    public create = async (request: Request, response: ReplyNoContinue): Promise<any> => {
        try {
            const todo = await this.repository.save(request.payload as ITodoEdit);

            return response({
                status: 201,
                data: todo,
            }).code(201);
        } catch (e) {
            throw Boom.badImplementation('Unable to create entity');
        }
    }
}
