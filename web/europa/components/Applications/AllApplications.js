import React, { useState } from 'react';

import EditApplication from './EditApplication';

const allApplications = ({
  updateClient,
  clients = [],
  loading = []
}) => {
  const [editingId, setEditingId] = useState(null);

  if (loading.includes('clients')) {
    return (
      <div className="applications-listing">
        <p className="text-center">Loading...</p>
      </div>
    );
  }

  if (!clients.length) {
    return (
      <div className="applications-listing">
        <p className="blank-list">No applications available yet.</p>
      </div>
    );
  }

  const applicationElements = clients.map((client, i) => (
    <div className="application-card" key={i}>
      <p className="title">
        {client.name}
        &nbsp;
        <small>(
          <a
            href={client.uri.split('\n')[0]}
            rel="noopener noreferrer"
            target="_blank">{client.uri.split('\n')[0].split(/:\/\//)[1]}
          </a>
        )</small>
      </p>
      <p className="description">{client.description}</p>
      <EditApplication
        updateClient={updateClient}
        cancelUpdate={() => setEditingId(null)}
        client={client}
        selectedClientId={editingId}
      />
      <ul className="inline-list all-applications-options">
        <li>
          <a
            href="#edit"
            onClick={(e) => {
              e.preventDefault();
              setEditingId(client.id);
            }}
          >
            Edit
          </a>
        </li>
        <li>
          <a
            href={`/api/clients/${client.id}/credentials`}
            title="It regenerates the client's secret for security reasons"
            rel="noopener noreferrer"
            target="_blank"
          >
            Download credentials
          </a>
        </li>
      </ul>
    </div>
  ));

  return (
    <div className="applications-listing">
      {applicationElements}
    </div>
  );
};

export default allApplications;
