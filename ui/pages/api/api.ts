import { NextApiRequest, NextApiResponse } from 'next'

export default function handler(req: NextApiRequest, res: NextApiResponse) {
  const { q } = req.query
  const results = {
    results: [
      { title: 'https://opentelemetry.io/', description: 'first-post' },
      { title: 'https://opentelemetry.io/', description: 'second-post' }
    ]
  }
  res.status(200).json(results)
}