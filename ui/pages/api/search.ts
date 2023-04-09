import { NextApiRequest, NextApiResponse } from 'next'

export default function handler(req: NextApiRequest, res: NextApiResponse) {
  const { q } = req.query
  const results = [
      { id: 1, title: 'https://opentelemetry.io/', description: 'first-post' },
      { id: 2, title: 'https://opentelemetry.io/', description: 'second-post' }
    ]
  
  res.status(200).json(results)
}