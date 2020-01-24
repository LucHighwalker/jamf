export default (defaultPort: number = 4200) => `import * as dotenv from 'dotenv';
dotenv.config();

import server from './server';

const port = process.env.PORT || ${defaultPort};

server.listen(port, (error: Error) => {
  if (error) {
    return console.log(error);
  }

  return console.log(\`server is listening on \${port}\`);
});
`;