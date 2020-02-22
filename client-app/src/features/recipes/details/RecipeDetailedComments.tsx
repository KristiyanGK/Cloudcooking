import React, { Fragment, useContext, useEffect, useState } from 'react';
import { Segment, Header, Form, Button, Comment } from 'semantic-ui-react';
import { RootStoreContext } from '../../../app/stores/rootStore';
import { observer } from 'mobx-react-lite';
import { Form as FinalForm, Field } from 'react-final-form';
import LoadingComponent from '../../../app/layout/LoadingComponent';
import TextInput from '../../../app/common/form/TextInput';
import { combineValidators, isRequired } from 'revalidate';

interface IProps {
  recipeId: string
}

const validate = combineValidators({
  content: isRequired({ message: 'The comment content is required' })
});

const RecipeDetailedComments: React.FC<IProps> = ({recipeId}) => {
  const rootStore = useContext(RootStoreContext);

  const { comments, loadingComments, loadComments, createComment, submitting } = rootStore.commentStore;

  const [loading, setLoading] = useState(false);

  useEffect(() => {
    setLoading(true);
    loadComments(recipeId).finally(() => setLoading(false));
  }, [loadComments, recipeId]);

  const handleFinalFormSubmit = (values: any) => {
    const { ...comment } = values;
    createComment(comment, recipeId)
  };

  if (loadingComments) return (<LoadingComponent content="Loading Comments..."/>)

  return (
    <Fragment>
      <Segment
        textAlign='center'
        attached='top'
        inverted
        color='teal'
        style={{ border: 'none' }}
      >
      <Header>Comments</Header>
      </Segment>
      <Segment attached>
        <Comment.Group>
          {comments && comments.map((comment) => (
            <Comment key={comment.id}>
            <Comment.Avatar src='/assets/user.png' />
            <Comment.Content>
              <Comment.Author>{comment.user}</Comment.Author>
              <Comment.Metadata>
                <div>Today at 5:43PM</div>
              </Comment.Metadata>
              <Comment.Text>{comment.content}</Comment.Text>
            </Comment.Content>
          </Comment>
          ))}
        </Comment.Group>

        <FinalForm
          onSubmit={handleFinalFormSubmit}
          render={({ handleSubmit, invalid, pristine }) => (
          <Form reply onSubmit={handleSubmit} loading={loading}>
            <Field
                name='content'
                placeholder='Content goes here...'
                component={TextInput}
            />
            <Button
              content='Add Reply'
              loading={submitting}
              disabled={loading || invalid || pristine}
              labelPosition='left'
              icon='edit'
              primary
            />
          </Form>
          )}
          />
      </Segment>
    </Fragment>
  );
};

export default observer(RecipeDetailedComments);
