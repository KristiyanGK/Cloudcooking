import React, { Fragment, useContext } from 'react'
import { Form, Button, Image, Grid } from 'semantic-ui-react'
import { Form as FinalForm, Field } from 'react-final-form';
import TextInput from '../../app/common/form/TextInput';
import { observer } from 'mobx-react-lite';
import { RootStoreContext } from '../../app/stores/rootStore';

const Chat = () => {
    const rootStore = useContext(RootStoreContext);

    const { addMessage, messages } = rootStore.chatStore;
    const { user } = rootStore.userStore;

    const handleFinalFormSubmit = (values: any) => {
        const { ...message } = values;
        message.sender = user?.username;
        addMessage(message);
    }

    return (
        <Fragment>
            <Grid>
                {messages.map((message) => (
                    <Grid.Row key={message.id}>
                        <Grid.Column width={1}>
                            <Image src='/assets/user.png' />
                            {message.sender}
                        </Grid.Column>
                        <Grid.Column width={15}>
                            <p>{message.body}</p>
                        </Grid.Column>
                    </Grid.Row>
                ))}
            </Grid>
            <br/>
            <FinalForm 
            onSubmit={handleFinalFormSubmit}
            render={({ handleSubmit, invalid, pristine, form }) => (
            <Form onSubmit={() => {handleSubmit(); form.reset()}}>
                <Field
                  name='body'
                  placeholder='Message'
                  component={TextInput}
                />
                <Button
                  disabled={invalid || pristine}
                  floated='right'
                  positive
                  type='submit'
                  content='Send'
                />
            </Form>
            )}/>
        </Fragment>
    )
}

export default observer(Chat)