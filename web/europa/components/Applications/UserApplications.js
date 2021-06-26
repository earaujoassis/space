import React from 'react';

const userApplications = ({
  revokeActiveClient,
  clients = [],
  loading = []
}) => {
  if (loading.includes('users')) {
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

  const applications = clients.map((client, i) => (
    <div className="application-card" key={i}>
      <p className="title">
        {client.name}
        &nbsp;
        <small>
          (
            <a
              href={client.uri.split('\n')[0]}
              rel="noopener noreferrer"
              target="_blank"
            >
              {client.uri.split('\n')[0].split(/:\/\//)[1]}
            </a>
          )
        </small>
      </p>
      <p className="description">{client.description}</p>
      <ul className="inline-list all-applications-options">
        <li>
          <a
            href="#revoke"
            onClick={(e) => {
              e.preventDefault();
              revokeActiveClient(client.id);
            }}
          >
            Revoke access
          </a>
        </li>
      </ul>
    </div>
  ));

  return (
    <div className="applications-listing">
      {applications}
    </div>
  );
};

export default userApplications;
