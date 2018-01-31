import { inject, injectable } from 'inversify';
import CrudController from 'app/common/crud-controller';
import { IArticle, IArticleResolver, IUsersController } from 'app/interfaces';
import Types from 'app/ioc/types';
import * as Boom from "boom";
import * as Hapi from "hapi";
import * as grpc from 'grpc';
import * as todoMessages from 'app/protobuf/todo_pb';
import * as todoService from 'app/protobuf/todo_grpc_pb';
import * as fs from 'fs';
import {ILogger} from "../../interfaces";

// const PROTO_PATH = process.cwd() + '/src/protobuf/todo.proto';
// const protoDescriptor = grpc.load({ file: PROTO_PATH, root: process.cwd()});
// const todo = grpc.load(PROTO_PATH).todo;
// const client = new todo.Todo('localhost:8081', grpc.credentials.createInsecure());
// const root = fs.readFileSync('/run/secrets/server_crt');
// const cert = fs.readFileSync('/run/secrets/server_crt');
// const key = fs.readFileSync('/run/secrets/server_key');
// const client = new todoService.TodoClient('localhost:8081', grpc.credentials.createSsl(undefined, key, cert));
const client = new todoService.TodoClient('grpc-api:443', grpc.credentials.createInsecure());
@injectable()
export default class ArticlesController extends CrudController<IArticle> implements IUsersController {
    protected logger: ILogger;

    constructor(@inject(Types.Resolvers.ArticleResolver) resolver: IArticleResolver, @inject(Types.Services.Logger) logger: ILogger) {
        super(resolver);
        this.logger = logger;
    }

    public getAll = async (request: Hapi.Request, response: Hapi.ReplyNoContinue): Promise<any> => {
        const start = new Date().getMilliseconds()
        const filter = new todoMessages.TodoFilter();
        filter.setId('')
        const call = client.getTodos(filter)
        call.on('data', (todo: todoMessages.TodoResponse) => {
            console.log(todo.getCreatedat().)
            this.logger.info('found todo: ' + todo.getName())
        })
        call.on('end', () => {
            const time = new Date().getMilliseconds() - start;
            this.logger.info(`Finished in ${time}ms`)
        })
        call.on('status', () => this.logger.info('status'))
        // const update = new todoMessages.TodoUpdateRequest()
        // update.setId("5a7124b8f092db0001c8a923")
        // update.setComplete(false)
        // update.setName(update.getName())
        // client.updateTodo(update, (err, resp) => {
        //     this.logger.info(resp.getName())
        // })
        try {
            const entities: IArticle[] = await this.crudResolver.getAll();

            return response({
                statusCode: 200,
                data: entities,
            });
        } catch (error) {
            return response(Boom.badImplementation(error));
        }
    };
}
