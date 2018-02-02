import { inject, injectable, interfaces } from 'inversify';
import Types from '../../ioc/types';
import {ITodo, ITodoEdit, ITodoRepository} from "../../interfaces";
import {ITodoClient} from "../../protobuf/todo_grpc_pb";
import {TodoFilter, TodoResponse, TodoUpdateRequest} from "../../protobuf/todo_pb";
import {TodoCreateRequest} from "../../protobuf/todo_pb";

@injectable()
class TodoRepository implements ITodoRepository {
    /**
     * gRPC client instance
     */
    protected todos: ITodoClient;

    // public getById: (_id: string) => Promise<ITodo>;
    // public save: (data: ITodo) => Promise<ITodo>;
    // public updateById: (_id: string, data: ITodo) => Promise<ITodo>;
    public deleteById: (_id: string) => Promise<string>;

    /**
     * @param {interfaces.Factory<INedbDatastore>} factory
     */
    public constructor(@inject(Types.Services.Todos) todos: ITodoClient) {
        this.todos = todos;
    }

    public getAll(): Promise<ITodo[]> {
        return new Promise((resolve, reject) => {
            const todos: ITodo[] = [];
            const call = this.todos.getTodos(this.createFilter(''));

            call.on('data', (todo: TodoResponse) => {
                todos.push({
                    id: todo.getId(),
                    name: todo.getName(),
                    completed: todo.getComplete(),
                    createdAt: new Date(todo.getCreatedat() * 1000),
                    updatedAt: new Date(todo.getUpdatedat() * 1000),
                });
            });

            call.on('error', (err) => {
                reject(err);
            });

            call.on('end', () => resolve(todos));
        });
    }

    public getById(id: string): Promise<ITodo> {
        return new Promise((resolve, reject) => {
            const todos: ITodo[] = [];
            const call = this.todos.getTodos(this.createFilter(id));

            call.on('data', (todo: TodoResponse) => {
                todos.push({
                    id: todo.getId(),
                    name: todo.getName(),
                    completed: todo.getComplete(),
                    createdAt: new Date(todo.getCreatedat() * 1000),
                    updatedAt: new Date(todo.getUpdatedat() * 1000),
                });
            });

            call.on('error', (err) => {
                reject(err);
            });

            call.on('end', () => resolve(todos[0]));
        });
    }

    public updateById(id: string, data: ITodoEdit): Promise<ITodo> {
        return new Promise((resolve, reject) => {
            if (!data.completed || !data.name) {
                reject(new Error('Update request must contain name and completed status'));
            }
            const completed = data.completed;
            const update = new TodoUpdateRequest();
            update.setId(id);
            update.setComplete(completed);
            update.setName(data.name);

            this.todos.updateTodo(update, (err, resp) => {
                if (err) {
                    reject(err);
                }

                resolve(resp.toObject());
            });
        });
    }

    public async save(data: ITodoEdit): Promise<ITodo> {
        return new Promise((resolve, reject) => {
            const completed = data.completed;

            const update = new TodoCreateRequest();
            update.setComplete(completed);
            update.setName(data.name);

            this.todos.createTodo(update, (err, todo) => {
                if (err) {
                    reject(err);
                }

                resolve({
                    id: todo.getId(),
                    name: todo.getName(),
                    completed: todo.getComplete(),
                    createdAt: new Date(todo.getCreatedat() * 1000),
                    updatedAt: new Date(todo.getUpdatedat() * 1000),
                });
            });
        });
    }

    protected createFilter(id: string): TodoFilter {
        const filter = new TodoFilter();
        filter.setId(id);

        return filter;
    }
}

export default TodoRepository;
