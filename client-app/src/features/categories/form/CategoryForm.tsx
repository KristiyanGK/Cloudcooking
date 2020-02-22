import React, { useState, useContext, useEffect } from 'react';
import { Segment, Form, Button, Grid } from 'semantic-ui-react';
import { observer } from 'mobx-react-lite';
import { RouteComponentProps } from 'react-router';
import { Form as FinalForm, Field } from 'react-final-form';
import TextInput from '../../../app/common/form/TextInput';
import TextAreaInput from '../../../app/common/form/TextAreaInput';
import SelectInput from '../../../app/common/form/SelectInput';
import { category } from '../../../app/common/options/categoryOptions';
import {
  combineValidators,
  isRequired,
  composeValidators,
  hasLengthGreaterThan
} from 'revalidate';
import { RootStoreContext } from '../../../app/stores/rootStore';
import { RecipeFormValues } from '../../../app/models/recipe';

const validate = combineValidators({
  title: isRequired({ message: 'The recipe title is required' }),
  category: isRequired('Category'),
  description: composeValidators(
    isRequired('Description'),
    hasLengthGreaterThan(4)({
      message: 'Description needs to be at least 5 characters'
    })
  )(),
  usedProducts: isRequired('UsedProducts'),
  cookingTime: isRequired('CookingTime')
});

interface DetailParams {
  id: string;
}

const CategoryForm: React.FC<RouteComponentProps<DetailParams>> = ({
  match,
  history
}) => {
  const rootStore = useContext(RootStoreContext);
  const {
    createRecipe,
    editRecipe,
    submitting,
    loadRecipe
  } = rootStore.recipeStore;

  const [recipe, setRecipe] = useState(new RecipeFormValues());
  const [loading, setLoading] = useState(false);

  useEffect(() => {
    if (match.params.id) {
      setLoading(true);
      loadRecipe(match.params.id)
        .then(recipe => {
          setRecipe(new RecipeFormValues(recipe));
        })
      .finally(() => setLoading(false));
    }
  }, [loadRecipe, match.params.id]);

  const handleFinalFormSubmit = (values: any) => {
    const { ...recipe } = values;
    if (!recipe.id) {
      let newRecipe = {
        ...recipe
      };
      createRecipe(newRecipe);
    } else {
      editRecipe(recipe);
    }
  };

  return (
    <Grid>
      <Grid.Column width={10}>
        <Segment clearing>
          <FinalForm
            validate={validate}
            initialValues={recipe}
            onSubmit={handleFinalFormSubmit}
            render={({ handleSubmit, invalid, pristine }) => (
              <Form onSubmit={handleSubmit} loading={loading}>
                <Field
                  name='title'
                  placeholder='Title'
                  value={recipe.title}
                  component={TextInput}
                />
                <Field
                  name='description'
                  placeholder='Description'
                  rows={3}
                  value={recipe.description}
                  component={TextAreaInput}
                />
                <Field
                  component={SelectInput}
                  options={category}
                  name='category'
                  placeholder='Category'
                  value={recipe.category?.name}
                />
                <Field
                  component={TextInput}
                  name='cookingTime'
                  type='number'
                  placeholder={5}
                  value={recipe.cookingTime + ""}
                />
                <Field
                  component={TextInput}
                  name='used Products'
                  placeholder='usedProducts'
                  value={recipe.usedProducts}
                />
                <Button
                  loading={submitting}
                  disabled={loading || invalid || pristine}
                  floated='right'
                  positive
                  type='submit'
                  content='Submit'
                />
                <Button
                  onClick={
                    recipe.id
                      ? () => history.push(`/recipes/${recipe.id}`)
                      : () => history.push('/recipes')
                  }
                  disabled={loading}
                  floated='right'
                  type='button'
                  content='Cancel'
                />
              </Form>
            )}
          />
        </Segment>
      </Grid.Column>
    </Grid>
  );
};

export default observer(CategoryForm);
