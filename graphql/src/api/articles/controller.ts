import { inject, injectable } from 'inversify';
import CrudController from 'app/common/crud-controller';
import { IArticle, IArticleResolver, IArticlesController } from 'app/interfaces';
import Types from 'app/ioc/types';

@injectable()
export default class ArticlesController extends CrudController<IArticle> implements IArticlesController {
    constructor(
        @inject(Types.Resolvers.ArticleResolver) resolver: IArticleResolver
    ) {
        super(resolver);
    }
}
