'use client';

import React from 'react';
import { useKeycloak } from '../context/KeycloakContext';
import { createServer } from '../lib/api/apiClient';
import { Formik, Form, Field, ErrorMessage } from 'formik';
import * as Yup from 'yup';

export default function CreateServerForm({ onCreated }) {
    const { keycloak, authenticated } = useKeycloak();

    const servNameMinLen = 3;
    const servNameMaxLen = 20;

    const createServerSchema = Yup.object().shape({
        name: Yup.string()
            .trim()
            .required('Server name is required')
            .min(servNameMinLen, `Server name has to be at least ${servNameMinLen} characters long`)
            .max(servNameMaxLen, `Server name can't be longer than ${servNameMaxLen} characters`),
    });

    const handleSubmit = async (values, { setSubmitting, resetForm, setStatus }) => {
        setStatus(null);

        try {
            const token = keycloak.token;
            const data = await createServer(token, values.name.trim());
            resetForm();
            onCreated(data.server);
        } catch (err) {
            console.error(err);
            setStatus('Could not create server');
        } finally {
            setSubmitting(false);
        }
    };

    if (!authenticated) {
        return;
    }

    return (
        <div className="mb-4">
            <Formik
                initialValues={{ name: '' }}
                validationSchema={createServerSchema}
                onSubmit={handleSubmit}
            >
                {({ isSubmitting, status }) => (
                    <Form>
                        <div className="mb-3">
                            <label>Server Name</label>
                            <Field
                                name="name"
                                type="text"
                                className="form-control"
                                placeholder="Server Name"
                                disabled={isSubmitting}
                            />
                            <ErrorMessage 
                                name="name" 
                                component="div" 
                                className="text-danger small"
                            />
                        </div>

                        {status && (
                            <div className="alert alert-danger" role="alert">
                                {status}
                            </div>
                        )}

                        <button
                            type="submit"
                            className="btn btn-primary"
                            disabled={isSubmitting}
                        >
                        {isSubmitting ? (
                            <>
                                <span
                                    className="spinner-border spinner-border-sm me-1"
                                    role="status"
                                    aria-hidden="true"
                                ></span>
                                Creating the server...
                            </>
                        ) : (
                            'Create Server'
                        )}
                        </button>
                    </Form>
                )}
            </Formik>
        </div>
    )
}
