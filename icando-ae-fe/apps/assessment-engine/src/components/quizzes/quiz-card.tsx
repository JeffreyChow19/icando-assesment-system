import {
    Card,
    CardContent,
    CardDescription,
    CardFooter,
    CardHeader,
    CardTitle,
} from "@ui/components/ui/card.tsx"

import { useEffect, useState } from 'react';
import { Button } from '@ui/components/ui/button.tsx';

export function QuizCard() {
    return (
        <Card className="space-x-2">
            <CardHeader className="flex flex-row justify-between">
                <CardTitle>Quiz Title</CardTitle>
                <p className="text-green-500">Published</p>
            </CardHeader>
            <CardContent>
                <CardDescription >
                    <p>Quiz Subject</p>
                </CardDescription>
                <CardDescription >
                    <p>2 Questions</p>
                </CardDescription>
            </CardContent>
            <CardFooter className="flex justify-between">
                <div className="flex flex-column justify-between">
                    <div>
                        <p>Last Published at: -</p>
                        <p>Updated at: 2021-09-20 12:00</p>
                    </div>
                </div>
                <div className="flex flex-row justify-between space-x-2">
                    <Button variant="outline">Edit</Button>
                    <Button>Publish</Button>
                </div>
            </CardFooter>
        </Card>
    )
}