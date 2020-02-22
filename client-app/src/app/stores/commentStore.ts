import { RootStore } from "./rootStore";
import { observable, action, reaction, runInAction } from "mobx";
import { IComment } from "../models/comment";
import agent from "../api/agent";
import { toast } from "react-toastify";

export default class CommentStore {
    rootStore: RootStore;

    constructor(rootStore: RootStore) {
        this.rootStore = rootStore
    }

    @observable comments: IComment[] = [];
    @observable loadingComments = false;
    @observable submitting = false;

    @action loadComments = async(recipeId: string) => {
        this.comments = [];
        this.loadingComments = true;
        try {
          const comments = await agent.Comments.list(recipeId);
          runInAction('loading comments', () => {
            comments.forEach(comment => {
              this.comments.push(comment);
            });
            this.loadingComments = false;
          })
        } catch (error) {
          runInAction('load comments error', () => {
            this.loadingComments = false;
          })
        }
    }

    @action createComment = async(comment: IComment, recipeId: string) => {
        this.submitting = true;
        try {
          const commentResult = await agent.Comments.create(comment, recipeId);
          runInAction('create comment', () => {
            this.comments.push(commentResult);
            this.submitting = false;
          })
        } catch (error) {
          runInAction('create comment error', () => {
            this.submitting = false;
          })
          toast.error('Problem submitting data');
          console.log(error.response);
        }
    }
}